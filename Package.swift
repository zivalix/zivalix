// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "zivalix",
    products: [
        .library(name: "zivalix", targets: ["zivalix"]),
    ],
    targets: [
        .target(name: "zivalix", path: "Sources", sources: ["zivalix/zivalix.swift"]),
    ],
    swiftLanguageVersions: [.v5]
)
