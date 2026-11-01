Pod::Spec.new do |s|
  s.name         = "zivalix"
  s.version      = "0.1.0"
  s.summary      = "Official reserved package for ZIVALIX Agency."
  s.description  = "Official reserved package for ZIVALIX Agency. Defensive registration to prevent dependency confusion attacks."
  s.homepage     = "https://zivalix.com"
  s.license      = { :type => "ISC" }
  s.author       = { "ZIVALIX Founders" => "founders@zivalix.com" }
  s.source       = { :git => "https://github.com/zivalix/zivalix.git", :commit => "YOUR_COMMIT_SHA" }
  s.source_files = "Sources/**/*.swift"
  s.swift_version = "5.9"
  s.static_framework = true
end
