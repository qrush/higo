chmod = `chmod -v 755 higo`
print chmod
puts "LS SOMETHING"
puts
dir = `ls -lah`
print dir

port = ARGV[0]
puts
puts "CUR DIR"
dir = `ls -lah ./`
print dir
puts
system "./higo -port #{port}"
