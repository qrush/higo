task :run do
  if ENV['PORT']
    sh "./higo -port #{ENV['PORT']}"
  else
    sh "./higo"
  end
end

task :clean do
  sh "make clean"
  sh "cd vendor/mango && make clean"
end

task :build do
  sh "cd vendor/mango && make"
  sh "mv vendor/mango/_obj/mango.a ."
  sh "make"
end

task :default => [:clean, :build, :run]
