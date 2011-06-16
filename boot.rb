pwd  = `pwd`.strip
chmod = `chmod -v 755 higo`
print chmod
dir = `ls -lah`
print dir

port = ARGV[0]
puts "starting this up!"
puts "pwd: #{pwd}"
puts "port: #{port}"
exec "#{File.join(pwd, "higo")} -port #{port}"
