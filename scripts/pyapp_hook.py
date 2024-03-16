# ruff: noqa: ARG002
from __future__ import annotations

import os
import sys
from typing import Any

from hatchling.builders.hooks.plugin.interface import BuildHookInterface


class _Config:
    def __init__(
        self,
        *,
        pyapp_version: str | None,
        requirements: str | None,
        features: list[str],
        exec_module: str | None,
        exec_spec: str | None,
        exec_code: str | None,
        python_version: str,
        python_variant: str | None,
        pip_external: bool | None,
        pip_version: str | None,
        pip_extra_args: str | None,
        pip_allow_config: bool | None,
        full_isolation: bool | None,
        upgrade_virtualenv: bool | None,
        pass_location: bool | None,
        self_command: str | None,
    ):
        self.pyapp_version: str | None = pyapp_version
        self.requirements: str | None = requirements
        self.features: list[str] = features

        self.exec_module: str | None = exec_module
        self.exec_spec: str | None = exec_spec
        self.exec_code: str | None = exec_code

        self.python_version: str = python_version
        self.python_variant: str | None = python_variant

        self.pip_external: bool | None = pip_external
        self.pip_version: str | None = pip_version
        self.pip_extra_args: str | None = pip_extra_args
        self.pip_allow_config: bool | None = pip_allow_config

        self.full_isolation: bool | None = full_isolation
        self.upgrade_virtualenv: bool | None = upgrade_virtualenv
        self.pass_location: bool | None = pass_location
        self.self_command: str | None = self_command


def _get_setting(
    config: dict[str, Any],
    env: os._Environ[str],
    env_var: str,
    setting_var: str,
    default: Any,
):
    return env.get(env_var, config.get(setting_var, default))


def _str2bool(value: str | bool | None) -> bool | None:
    if isinstance(value, bool):
        return value
    if value is None:
        return None
    return value in ("true", "1")


def _parse_config(plugin_name: str, config: dict[str, Any], env: os._Environ[str]):
    pyapp_version = config.get("pyapp-version", None)

    req_file = _get_setting(
        config, env, "PYAPP_PROJECT_DEPENDENCY_FILE", "requirements", None
    )
    features = _get_setting(config, env, "PYAPP_PROJECT_FEATURES", "features", [])
    if isinstance(features, str):
        features = features.split(",")

    exec_module = _get_setting(config, env, "PYAPP_EXEC_MODULE", "exec-module", None)
    exec_spec = _get_setting(config, env, "PYAPP_EXEC_SPEC", "exec-spec", None)
    exec_code = _get_setting(config, env, "PYAPP_EXEC_CODE", "exec-code", None)
    if exec_module is None and exec_spec is None and exec_code is None:
        required_fields = "'exec_module','exec_spec','exec_code'"
        message = f"`{plugin_name}` missing one of required fields {required_fields}"
        raise TypeError(message)

    python_version = _get_setting(
        config, env, "PYAPP_PYTHON_VERSION", "python-version", None
    )
    python_variant = _get_setting(
        config, env, "PYAPP_DISTRIBUTION_VARIANT", "python-variant", None
    )
    if not isinstance(python_version, str):
        message = f"`{plugin_name}` 'python-version' is required and must be a string"
        raise TypeError(message)

    pip_external = _str2bool(
        _get_setting(config, env, "PYAPP_PIP_EXTERNAL", "pip-external", None)
    )
    pip_version = _get_setting(config, env, "PYAPP_PIP_VERSION", "pip-version", None)
    pip_extra_args = _get_setting(
        config, env, "PYAPP_PIP_EXTRA_ARGS", "pip-extra-args", None
    )
    pip_allow_config = _str2bool(
        _get_setting(config, env, "PYAPP_PIP_ALLOW_CONFIG", "pip-allow-config", None)
    )

    full_isolation = _str2bool(
        _get_setting(config, env, "PYAPP_FULL_ISOLATION", "full-isolation", None)
    )
    upgrade_virtualenv = _str2bool(
        _get_setting(config, env, "PYAPP_UPGRADE_VIRTUALENV", "virtualenv", None)
    )
    pass_location = _str2bool(
        _get_setting(config, env, "PYAPP_PASS_LOCATION", "pass-location", None)
    )
    self_command = _get_setting(config, env, "PYAPP_SELF_COMMAND", "self-command", None)

    return _Config(
        pyapp_version=pyapp_version,
        requirements=req_file,
        features=features,
        exec_module=exec_module,
        exec_spec=exec_spec,
        exec_code=exec_code,
        python_version=python_version,
        python_variant=python_variant,
        pip_external=pip_external,
        pip_version=pip_version,
        pip_extra_args=pip_extra_args,
        pip_allow_config=pip_allow_config,
        full_isolation=full_isolation,
        upgrade_virtualenv=upgrade_virtualenv,
        pass_location=pass_location,
        self_command=self_command,
    )


def _config2env(plugin_name: str, config: _Config, env: dict[str, str]):
    if config.requirements is not None:
        env["PYAPP_PROJECT_DEPENDENCY_FILE"] = os.path.abspath(config.requirements)

    if len(config.features) > 0:
        env["PYAPP_PROJECT_FEATURES"] = ",".join(config.features)

    if config.exec_module is not None:
        env["PYAPP_EXEC_MODULE"] = config.exec_module
    elif config.exec_spec is not None:
        env["PYAPP_EXEC_SPEC"] = config.exec_spec
    elif config.exec_code is not None:
        env["PYAPP_EXEC_CODE"] = config.exec_code
    else:
        required_fields = "'exec_module','exec_spec','exec_code'"
        message = f"`{plugin_name}` missing one of required fields {required_fields}"
        raise TypeError(message)

    env["PYAPP_PYTHON_VERSION"] = config.python_version
    if config.python_variant is not None:
        env["PYAPP_DISTRIBUTION_VARIANT"] = config.python_variant

    if config.pip_external is True:
        env["PYAPP_PIP_EXTERNAL"] = "true"
    if config.pip_version is not None:
        env["PYAPP_PIP_VERSION"] = config.pip_version
    if config.pip_extra_args is not None:
        env["PYAPP_PIP_EXTRA_ARGS"] = config.pip_extra_args
    if config.pip_allow_config is True:
        env["PYAPP_PIP_ALLOW_CONFIG"] = "true"

    if config.full_isolation is True:
        env["PYAPP_FULL_ISOLATION"] = "true"
    if config.upgrade_virtualenv is True:
        env["PYAPP_UPGRADE_VIRTUALENV"] = "true"
    if config.pass_location is True:
        env["PYAPP_PASS_LOCATION"] = "true"  # noqa: S105
    if config.self_command is not None:
        env["PYAPP_SELF_COMMAND"] = config.self_command


class BinaryBuilderHook(BuildHookInterface):
    SUPPORTED_VERSIONS = ("3.12", "3.11", "3.10", "3.9", "3.8", "3.7")

    PLUGIN_NAME = "pyapp_hook"

    def finalize(
        self, version: str, build_data: dict[str, Any], artifact_path: str
    ) -> None:
        if self.target_name != "wheel":
            message = f"Hook '{self.plugin_name}' only works with the target 'wheel'"
            raise ValueError(message)

        config = _parse_config(self.PLUGIN_NAME, self.config, os.environ)

        import shutil
        import tempfile

        cargo_path = os.environ.get("CARGO", "")
        if not cargo_path:
            if not shutil.which("cargo"):
                message = "Executable `cargo` could not be found on PATH"
                raise OSError(message)

            cargo_path = "cargo"

        app_dir = self.directory
        if not os.path.isdir(app_dir):
            os.makedirs(app_dir)

        on_windows = sys.platform == "win32"
        base_env = dict(os.environ)
        base_env["PYAPP_PROJECT_PATH"] = artifact_path
        # The variables PYAPP_PROJECT_NAME, PYAPP_PROJECT_VERSION are derived from
        # PYAPP_PROJECT_PATH: https://ofek.dev/pyapp/latest/config/#project-embedding
        _config2env(self.PLUGIN_NAME, config, base_env)
        # TODO: Add support for custom and embeded distribution
        base_env.pop("PYAPP_EXEC_SCRIPT", None)
        base_env.pop("PYAPP_EXEC_NOTEBOOK", None)
        base_env.pop("PYAPP_SKIP_INSTALL", None)

        # https://doc.rust-lang.org/cargo/reference/config.html#buildtarget
        build_target = os.environ.get("CARGO_BUILD_TARGET", "")

        # This will determine whether we install from crates.io or build locally and is
        # currently required for cross compilation:
        # https://github.com/cross-rs/cross/issues/1215
        repo_path = os.environ.get("PYAPP_REPO", "")

        with tempfile.TemporaryDirectory() as temp_dir:
            exe_name = "pyapp.exe" if on_windows else "pyapp"
            if repo_path:
                context_dir = repo_path
                target_dir = os.path.join(temp_dir, "build")
                if build_target:
                    temp_exe_path = os.path.join(
                        target_dir, build_target, "release", exe_name
                    )
                else:
                    temp_exe_path = os.path.join(target_dir, "release", exe_name)
                install_command = [
                    cargo_path,
                    "build",
                    "--release",
                    "--target-dir",
                    target_dir,
                ]
            else:
                context_dir = temp_dir
                temp_exe_path = os.path.join(temp_dir, "bin", exe_name)
                install_command = [
                    cargo_path,
                    "install",
                    "pyapp",
                    "--force",
                    "--root",
                    temp_dir,
                ]
                if config.pyapp_version:
                    install_command.extend(["--version", config.pyapp_version])

            self.cargo_build(install_command, cwd=context_dir, env=base_env)

            exe_stem = (
                f"{self.metadata.name}-{self.metadata.version}-{build_target}"
                if build_target
                else f"{self.metadata.name}-{self.metadata.version}"
            )
            exe_path = os.path.join(
                app_dir, f"{exe_stem}.exe" if on_windows else exe_stem
            )
            shutil.move(temp_exe_path, exe_path)

    def cargo_build(self, *args: Any, **kwargs: Any) -> None:
        import subprocess

        if self.app.verbosity < 0:
            kwargs["stdout"] = subprocess.PIPE
            kwargs["stderr"] = subprocess.STDOUT

        process = subprocess.run(*args, **kwargs)  # noqa: PLW1510
        if process.returncode:
            message = f"Compilation failed (code {process.returncode})"
            raise OSError(message)
