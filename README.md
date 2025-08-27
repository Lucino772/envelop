# envelop :construction:

> [!WARNING]
> **`envelop`** is currently in its early stages of development. Features, APIs, and functionality are subject to change as the project evolves.

**`envelop`** is a lightweight tool designed to install, run, and manage dedicated game servers. It provides a unified management interface that abstracts common protocols like **RCON** and **Query**, delivering a consistent **API** across different game servers. The interface includes essential services such as **player management**, **process control** (status, logs, stdin input), **RCON command execution**, and **event streaming** (e.g., server status changes, player activity).

Additionally, **`envelop`** simplifies server installation, supporting both **SteamCMD-compatible games** and **non-Steam game servers**, aiming to be the go-to solution for deploying and managing any game server seamlessly.

## 🛠️ Steam Integration
The **Steam** integration in **`envelop`** was made possible thanks to the amazing work of the **[SteamKit2](https://github.com/SteamRE/SteamKit)** project. In addition, the **`steamvdf`** package, which handles parsing Valve Data Format files, was adapted from **[Jleagle/steam-go](https://github.com/Jleagle/steam-go)**.