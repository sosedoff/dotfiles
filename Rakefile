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

desc 'Install rubies'
task :rubies do
  %(1.8.7 1.9.2 1.9.3 ree rbx).each do |r|
    puts "Installing #{r}"
    `rvm install #{r}`
  end
end

desc 'Install Sublime Text 2 settings'
task :sublime_text2 do
  source_path = File.join(ENV['HOME'], '.sublime_text2', 'Preferences.sublime-settings')
  target_path = File.expand_path("~/Library/Application Support/Sublime Text 2/Packages/User")

  if File.exists?(target_path)
    system %w[unlink #{target_path}]
  end

  Dir.chdir File.dirname(__FILE__) do
    system %[ln -vsf #{source_path} \"#{target_path}\"]
  end
end
