desc 'Install all dotfiles'
task :install do
  home = ENV['HOME']
  skip_files = %w(Rakefile README.md)

  Dir.chdir File.dirname(__FILE__) do
    dotfiles_dir = Dir.pwd.sub(home + '/', '')
    
    Dir['*'].each do |file|
      next if skip_files.include?(file)
      next unless File.extname(file).empty?
      
      target_name = file == 'bin' ? file : ".#{file}"
      target = File.join(home, target_name)
      
      if File.exists?(target)
        system %[unlink #{target}]
      end
      
      system %[ln -vsf #{File.join(dotfiles_dir, file)} #{target}]
    end
  end
end

desc 'Install default gems'
task :gems do
  puts "Installing default gems..."
  File.readlines('default_gems').each do |gem|
    gem_name = gem.strip
    puts " -> #{gem_name}"
    `gem install #{gem_name}`
  end
end

desk 'Install rubies'
task :rubies do
  %(1.8.7 1.9.2 1.9.3 ree rbx).each do |r|
    puts "Installing #{r}"
    `rvm install #{r}`
  end
end
