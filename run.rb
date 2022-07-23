require "fileutils"

def gomod(v)
  File.open("go.mod", "w") do |f|
    f.puts <<~"GOMOD"
      module github.com/nabetani/gobe.git/v1.#{v}
      go 1.#{v}
    GOMOD
  end
end

(11..18).each do |v|
  dir = "v1.%d" % v
  cmd = v==18 ? "go" : "go1.%d" % v
  FileUtils.mkdir_p(dir)
  Dir.chdir(dir) do
    unless dir=="v1.11"
      FileUtils.rm_f( ["main.go", dir])
      FileUtils.cp( Dir.glob("../v1.11/*.go"), ".")
      File.open( ".gitignore", "w" ){ |f| f.puts("*") }
      gomod(v)
    end
    puts %x(#{cmd} build -o #{dir} main.go 2>&1)
    sleep(0.1)
    puts %x(./#{dir} 10 2>&1)
    5.times do
      sleep(0.1)
      puts %x(./#{dir} 50000 2>&1)
    end
  end
end