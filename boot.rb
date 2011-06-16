chmod = `chmod -v 755 higo`
print chmod
puts "LS SOMETHING"
puts
dir = `ls -lah`
print dir

port = ARGV[0]
puts
puts "uname"
dir = `uname -a`
print dir
puts
system "./higo -port #{port}"
