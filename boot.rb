pwd = ENV['PWD']
port = ARGV[0]

puts "starting this up!"
puts "pwd: #{pwd}"
puts "port: #{port}"
exec "#{File.join(pwd, "higo")} -port #{port}"
