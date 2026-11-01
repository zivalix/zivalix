from conan import ConanFile

class ZivalixDefensiveContract(ConanFile):
    name = "zivalix"
    version = "0.1.0"
    description = "Official reserved package for ZIVALIX Agency. Defensive registration."
    homepage = "https://zivalix.com"
    license = "ISC"
    author = "ZIVALIX Founders <founders@zivalix.com>"
    settings = "os", "compiler", "build_type", "arch"
    package_type = "library"
    no_copy_source = True

    def build(self):
        raise Exception("🛡️ ZIVALIX SECURE NAMESPACE: Unauthorized access strictly prohibited. Build Terminated.")
