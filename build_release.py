#! /usr/bin/python3
from sys import argv
import os
import subprocess
from shutil import copy2, make_archive, rmtree

app_name = "dirsort"
app_version = ""
if len(argv) > 1:
    app_version = argv[1]

build_dir = "./build/"
build_tmp_dir = os.path.join(build_dir, app_name, "")

files_to_archieve = ["config.yaml"]

platforms = ["windows/amd64", "windows/386", "linux/amd64", "linux/386", "linux/arm64", "darwin/amd64", "darwin/arm64"]

os_names = {
    "windows": "win",
    "darwin": "macos"
}
architectures = {
    "amd64": "x64",
    "386": "x32",
}
compressors = {
    "linux": "gztar",
    "darwin": "gztar",
    "windows": "zip",
}

def compress(name: str, format: str) -> str:
    name = os.path.join(build_dir, name)
    return make_archive(name, format, root_dir=build_dir, base_dir=app_name)

def generateArchieveName(target_os: str, architecture: str, sep="_") -> str:
    arch = architectures.get(architecture) or architecture
    os = os_names.get(target_os) or target_os
    return sep.join([app_name, app_version, os, arch])

def build(goos: str, goarch: str):
    env = os.environ
    env["GOOS"] = goos
    env["GOARCH"] = goarch
    try:
        go_build = subprocess.Popen(["go", "build", "-o", build_tmp_dir, "."], env=env)
        go_build.wait()
    except Exception as e:
        print("Building error:", e)

if not os.path.exists(build_tmp_dir):
    os.makedirs(build_tmp_dir)

for file in files_to_archieve:
    copy2(file, build_tmp_dir)

for platform in platforms:
    target_os, arch = platform.split("/")
    print("building for {}...".format(platform), end=" ")
    build(target_os, arch)
    print("done.")
    builded_executable = os.path.join(build_tmp_dir, app_name)
    if target_os == "windows":
        builded_executable += ".exe"
    compress(generateArchieveName(target_os, arch), compressors.get(target_os))
    os.remove(builded_executable)

rmtree(build_tmp_dir)