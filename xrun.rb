require "fileutils"

def gomod(v, pre)
  File.open("go.mod", "w") do |f|
    f.puts <<~"GOMOD"
      module github.com/nabetani/gobe.git/#{pre}v1.#{v}
      go 1.#{v}
    GOMOD
  end
end

def govers
  case RUBY_PLATFORM
  when /arm64\-darwin/
    16..18
  when /x86_64\-darwin/
    11..18
  else
    raise "unexpected RUBY_PLATFORM value"
  end
end

govers.each do |v|
  pre = ARGV[0]
  dir = pre+"v1.%d" % v
  cmd = v==18 ? "go" : "go1.%d" % v
  num = ARGV[1]
  FileUtils.mkdir_p(dir)
  Dir.chdir(dir) do
    unless dir==pre+"v1.18"
      FileUtils.rm_f( ["main.go", dir])
      FileUtils.cp( Dir.glob("../#{pre}v1.18/*.go"), ".")
      File.open( ".gitignore", "w" ){ |f| f.puts("*") }
      gomod(v, pre)
    end
    puts %x(#{cmd} build -o #{dir} main.go 2>&1)
    puts( %x(./#{dir} 10 2>&1) )
    5.times do 
      puts( %x(./#{dir} #{num} 2>&1) )
    end
  end
end