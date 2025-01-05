package steamcm

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
	"slices"
	"time"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
	"github.com/Lucino772/envelop/pkg/steam/steampb"
	"github.com/ulikunitz/xz/lzma"
	"google.golang.org/protobuf/proto"
)

const (
	STEAM3_MANIFEST_MAGIC        uint32 = 0x16349781
	PROTOBUF_PAYLOAD_MAGIC       uint32 = 0x71F617D0
	PROTOBUF_METADATA_MAGIC      uint32 = 0x1F4812BE
	PROTOBUF_SIGNATURE_MAGIC     uint32 = 0x1B81B817
	PROTOBUF_ENDOFMANIFEST_MAGIC uint32 = 0x32C415AB
)

const (
	VZIP_HEADER     uint16 = 0x5A56
	VZIP_FOOTER     uint16 = 0x767A
	VZIP_VERSION    uint8  = 'a'
	VZIP_HEADER_LEN        = 7
	VZIP_FOOTER_LEN        = 10
)

type ChunkData struct {
	ChunkId         []byte
	Checksum        uint32
	Offset          uint64
	CompressedLen   uint32
	UncompressedLen uint32
}

type FileData struct {
	Filename     string
	FilenameHash []byte
	Chunks       []ChunkData
	Flags        steamlang.EDepotFileFlag
	TotalSize    uint64
	FileHash     []byte
	LinkTarget   string
}

type DepotManifest struct {
	FilenamesEncrypted    bool
	DepotId               uint32
	ManifestGid           uint64
	CreationTime          time.Time
	TotalUncompressedSize uint64
	TotalComressedSize    uint64
	EncryptedCRC          uint32
	Files                 []FileData
}

type DepotChunk struct {
	Chunk ChunkData
	Data  []byte
}

func NewDepotManifest(data []byte, depotKey []byte) (*DepotManifest, error) {
	buff := bufio.NewReader(bytes.NewReader(data))
	magicBytes, err := buff.Peek(4)
	if err != nil {
		return nil, err
	}
	magic := binary.LittleEndian.Uint32(magicBytes)

	var manifest *DepotManifest
	if magic == STEAM3_MANIFEST_MAGIC {
		manifest, err = parseBinaryManifest(buff)
		if err != nil {
			return nil, err
		}
	} else {
		manifest, err = parseProtobufManifest(buff)
		if err != nil {
			return nil, err
		}
	}

	if manifest.FilenamesEncrypted {
		decryptedFiles := make([]FileData, 0)
		for _, file := range manifest.Files {
			encFilename, err := base64.StdEncoding.DecodeString(file.Filename)
			if err != nil {
				return nil, err
			}

			filename, err := steam.AESDecrypt(depotKey, encFilename)
			if err != nil {
				return nil, err
			}
			file.Filename = string(filename[:len(filename)-1])
			decryptedFiles = append(decryptedFiles, file)
		}
		manifest.Files = decryptedFiles
	}

	return manifest, nil
}

func parseBinaryManifest(rd *bufio.Reader) (*DepotManifest, error) {
	var magic uint32
	if err := binary.Read(rd, binary.LittleEndian, &magic); err != nil {
		return nil, err
	}

	var metadata struct {
		Version               uint32
		DepotId               uint32
		ManifestGid           uint64
		CreationTime          uint32
		FilenameEncrypted     uint32
		TotalUncompressedSize uint64
		TotalCompressedSize   uint64
		ChunkCount            uint32
		FileEntryCount        uint32
		FileMappingSize       uint32
		EncryptedCRC          uint32
		DecryptedCRC          uint32
		Flags                 uint32
	}
	if err := binary.Read(rd, binary.LittleEndian, &metadata); err != nil {
		return nil, err
	}

	manifest := &DepotManifest{
		FilenamesEncrypted:    metadata.FilenameEncrypted != 0,
		DepotId:               metadata.DepotId,
		ManifestGid:           metadata.ManifestGid,
		CreationTime:          time.Unix(int64(metadata.CreationTime), 0),
		TotalUncompressedSize: metadata.TotalUncompressedSize,
		TotalComressedSize:    metadata.TotalCompressedSize,
		EncryptedCRC:          metadata.EncryptedCRC,
		Files:                 make([]FileData, 0),
	}

	for i := 0; i < int(metadata.FileMappingSize); i++ {
		filename, err := rd.ReadString(0)
		if err != nil {
			return nil, err
		}
		var (
			totalSize uint64
			flags     uint32
		)
		if err := binary.Read(rd, binary.LittleEndian, &totalSize); err != nil {
			return nil, err
		}
		if err := binary.Read(rd, binary.LittleEndian, &flags); err != nil {
			return nil, err
		}

		hashContent := make([]byte, 20)
		if _, err := io.ReadFull(rd, hashContent); err != nil {
			return nil, err
		}
		hashFilename := make([]byte, 20)
		if _, err := io.ReadFull(rd, hashFilename); err != nil {
			return nil, err
		}

		var numChunks uint32
		if err := binary.Read(rd, binary.LittleEndian, &numChunks); err != nil {
			return nil, err
		}

		file := FileData{
			Filename:     filename[:len(filename)-1],
			FilenameHash: hashFilename,
			Chunks:       make([]ChunkData, 0),
			Flags:        steamlang.EDepotFileFlag(flags),
			TotalSize:    totalSize,
			FileHash:     hashContent,
			LinkTarget:   "",
		}

		for j := 0; j < int(numChunks); j++ {
			chunkId := make([]byte, 20)
			if _, err := io.ReadFull(rd, chunkId); err != nil {
				return nil, err
			}
			var checksum uint32
			if err := binary.Read(rd, binary.LittleEndian, &checksum); err != nil {
				return nil, err
			}

			var chunk struct {
				offset           uint64
				decompressedSize uint32
				compressedSize   uint32
			}
			if err := binary.Read(rd, binary.LittleEndian, &chunk); err != nil {
				return nil, err
			}

			file.Chunks = append(
				file.Chunks,
				ChunkData{
					ChunkId:         chunkId,
					Checksum:        checksum,
					Offset:          chunk.offset,
					CompressedLen:   chunk.compressedSize,
					UncompressedLen: chunk.decompressedSize,
				},
			)
		}
		manifest.Files = append(manifest.Files, file)
	}
	return manifest, nil
}

func parseProtobufManifest(rd *bufio.Reader) (*DepotManifest, error) {
	var (
		contentPayload   *steampb.ContentManifestPayload
		contentMetadata  *steampb.ContentManifestMetadata
		contentSignature *steampb.ContentManifestSignature
	)

	for {
		var magic uint32
		if err := binary.Read(rd, binary.LittleEndian, &magic); err != nil {
			return nil, err
		}
		if magic == PROTOBUF_ENDOFMANIFEST_MAGIC {
			break
		}

		switch magic {
		case PROTOBUF_PAYLOAD_MAGIC:
			var payloadLen uint32
			if err := binary.Read(rd, binary.LittleEndian, &payloadLen); err != nil {
				return nil, err
			}
			payloadBuff := make([]byte, payloadLen)
			if _, err := io.ReadFull(rd, payloadBuff); err != nil {
				return nil, err
			}
			contentPayload = new(steampb.ContentManifestPayload)
			if err := proto.Unmarshal(payloadBuff, contentPayload); err != nil {
				return nil, err
			}
		case PROTOBUF_METADATA_MAGIC:
			var payloadLen uint32
			if err := binary.Read(rd, binary.LittleEndian, &payloadLen); err != nil {
				return nil, err
			}
			payloadBuff := make([]byte, payloadLen)
			if _, err := io.ReadFull(rd, payloadBuff); err != nil {
				return nil, err
			}
			contentMetadata = new(steampb.ContentManifestMetadata)
			if err := proto.Unmarshal(payloadBuff, contentMetadata); err != nil {
				return nil, err
			}
		case PROTOBUF_SIGNATURE_MAGIC:
			var payloadLen uint32
			if err := binary.Read(rd, binary.LittleEndian, &payloadLen); err != nil {
				return nil, err
			}
			payloadBuff := make([]byte, payloadLen)
			if _, err := io.ReadFull(rd, payloadBuff); err != nil {
				return nil, err
			}
			contentSignature = new(steampb.ContentManifestSignature)
			if err := proto.Unmarshal(payloadBuff, contentSignature); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("invalid header")
		}
	}

	if contentMetadata == nil || contentPayload == nil || contentSignature == nil {
		return nil, errors.New("invalid manifest")
	}

	manifest := &DepotManifest{
		FilenamesEncrypted:    contentMetadata.GetFilenamesEncrypted(),
		DepotId:               contentMetadata.GetDepotId(),
		ManifestGid:           contentMetadata.GetGidManifest(),
		CreationTime:          time.Unix(int64(contentMetadata.GetCreationTime()), 0),
		TotalUncompressedSize: contentMetadata.GetCbDiskOriginal(),
		TotalComressedSize:    contentMetadata.GetCbDiskCompressed(),
		EncryptedCRC:          contentMetadata.GetCrcEncrypted(),
		Files:                 make([]FileData, 0),
	}

	for _, mapping := range contentPayload.GetMappings() {
		// FIXME: Do we need to do something with the path when the filename is encrypted
		file := FileData{
			Filename:     mapping.GetFilename(),
			FilenameHash: mapping.GetShaFilename(),
			Chunks:       make([]ChunkData, 0),
			Flags:        steamlang.EDepotFileFlag(mapping.GetFlags()),
			TotalSize:    mapping.GetSize(),
			FileHash:     mapping.GetShaContent(),
			LinkTarget:   mapping.GetLinktarget(),
		}
		for _, chunk := range mapping.GetChunks() {
			file.Chunks = append(
				file.Chunks,
				ChunkData{
					ChunkId:         chunk.GetSha(),
					Checksum:        chunk.GetCrc(),
					Offset:          chunk.GetOffset(),
					CompressedLen:   chunk.GetCbCompressed(),
					UncompressedLen: chunk.GetCbOriginal(),
				},
			)
		}
		manifest.Files = append(manifest.Files, file)
	}
	return manifest, nil
}

func NewDepotChunk(chunk ChunkData, data []byte, depotKey []byte) (*DepotChunk, error) {
	decrypted, err := steam.AESDecrypt(depotKey, data)
	if err != nil {
		return nil, err
	}

	var processedData []byte
	if len(decrypted) > 1 && decrypted[0] == 'V' && decrypted[1] == 'Z' {
		_data, err := decompressLzma(decrypted)
		if err != nil {
			return nil, err
		}
		processedData = _data
	} else {
		_data, err := decompressZip(decrypted)
		if err != nil {
			return nil, err
		}
		processedData = _data
	}

	if adler32(processedData) != chunk.Checksum {
		return nil, errors.New("hash mismatch")
	}

	return &DepotChunk{
		Chunk: chunk,
		Data:  processedData,
	}, nil
}

func decompressZip(data []byte) ([]byte, error) {
	rd, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	if len(rd.File) == 0 {
		return nil, errors.New("missing file")
	}
	file := rd.File[0]
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	var buff bytes.Buffer
	if _, err := io.Copy(&buff, rc); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func decompressLzma(data []byte) ([]byte, error) {
	rd := bytes.NewReader(data)

	// Validate header
	var header uint16
	if err := binary.Read(rd, binary.LittleEndian, &header); err != nil {
		return nil, err
	}
	if header != VZIP_HEADER {
		return nil, errors.New("invalid VZip header")
	}

	var version uint8
	if err := binary.Read(rd, binary.LittleEndian, &version); err != nil {
		return nil, err
	}
	if version != VZIP_VERSION {
		return nil, errors.New("unsupported VZip version")
	}

	// Read CRC and properties
	var expectedCRC uint32
	if err := binary.Read(rd, binary.LittleEndian, &expectedCRC); err != nil {
		return nil, err
	}
	properties := make([]byte, 5)
	if _, err := io.ReadFull(rd, properties); err != nil {
		return nil, err
	}

	// Read compressed data
	compressedData := make([]byte, len(data)-VZIP_HEADER_LEN-VZIP_FOOTER_LEN-5)
	if _, err := io.ReadFull(rd, compressedData); err != nil {
		return nil, err
	}

	// Read footer
	var outputCRC uint32
	if err := binary.Read(rd, binary.LittleEndian, &outputCRC); err != nil {
		return nil, err
	}
	var sizeDecompressed uint32
	if err := binary.Read(rd, binary.LittleEndian, &sizeDecompressed); err != nil {
		return nil, err
	}
	var footer uint16
	if err := binary.Read(rd, binary.LittleEndian, &footer); err != nil {
		return nil, err
	}
	if footer != VZIP_FOOTER {
		return nil, errors.New("invalid VZip footer")
	}

	// Decompress data
	adjustedData := slices.Concat(
		properties,
		binary.LittleEndian.AppendUint64([]byte{}, uint64(sizeDecompressed)),
		compressedData,
	)

	lzmaReader, err := lzma.NewReader(bytes.NewReader(adjustedData))
	if err != nil {
		return nil, err
	}
	decompressedData := make([]byte, sizeDecompressed)
	if _, err := io.ReadFull(lzmaReader, decompressedData); err != nil {
		return nil, err
	}

	// Verify CRC
	actualCRC := crc32.ChecksumIEEE(decompressedData)
	if actualCRC != outputCRC {
		return nil, errors.New("CRC mismatch, data may be corrupted")
	}

	return decompressedData, nil
}

func adler32(data []byte) uint32 {
	var (
		a uint32
		b uint32
	)
	for i := 0; i < len(data); i++ {
		a = (a + uint32(data[i])) % 65521
		b = (b + a) % 65521
	}
	return a | (b << 16)
}
