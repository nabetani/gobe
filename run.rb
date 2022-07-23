require "fileutils"

[
  ["/usr/local/Cellar/go@1.13/1.13.15/bin/go", "v13"],
  ["/usr/local/Cellar/go@1.15/1.15.15/bin/go", "v15"],
  ["go", "v18"],
].each do |cmd,dir|
  Dir.chdir(dir) do
    unless dir=="v13"
      FileUtils.cp( Dir.glob("../v13/*.go"), ".")
    end
    puts(%x(#{cmd} run main.go))
  end
end