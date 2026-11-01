Pod::Spec.new do |s|
  s.name         = "zivalix"
  s.version      = "0.1.0"
  s.summary      = "Official reserved package for ZIVALIX Agency."
  s.description  = "Official reserved package for ZIVALIX Agency. Defensive registration to prevent dependency confusion attacks."
  s.homepage     = "https://zivalix.com"
  s.license      = { :type => "ISC" }
  s.author       = { "ZIVALIX Founders" => "founders@zivalix.com" }
  s.source       = { :git => "https://github.com/zivalix/zivalix.git", :commit => "139e6e461d7de5cab915d0ee04bcd91b2d523ac3" }
  s.source_files = "Sources/**/*.swift"
  s.swift_version = "5.9"
  s.static_framework = true
  s.platforms = { :ios => "13.0", :osx => "10.15" }
end
