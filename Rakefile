DOTFILES_ROOT = File.dirname(__FILE__)

desc 'Install all dotfiles'
task :dotfiles do
  puts "Installing DOT files..."
  home = ENV['HOME']
  dotfiles_dir = "#{File.dirname(__FILE__)}/dotfiles"
  Dir["#{dotfiles_dir}/*"].each do |file|
    link_path = File.join(home, ".#{File.basename(file)}")
    system %x[unlink #{link_path}] if File.exists?(link_path)
    system %[ln -vsf #{file} #{link_path}]
  end
end

desc 'Install all bin files'
task :binfiles do
  source_path = File.join(File.dirname(__FILE__), 'bin')
  target_path = File.join(ENV['HOME'], '.bin')

  if File.exists?(target_path)
    puts "Unlinking installed bin directory"
    system %[unlink #{target_path}]
  end

  puts "Installing new bin directory"
  puts "ln -vsf #{source_path} #{target_path}"
  system %[ln -vsf #{source_path} #{target_path}]
end

desc 'Install default gems'
task :gems do
  puts "Installing GEM files..."
  File.readlines('default_gems').each do |gem|
    gem_name = gem.strip
    puts " -> #{gem_name}"
    `gem install #{gem_name}`
  end
end

namespace :sublime do
  desc 'Install settings'
  task :settings do
    root_path   = File.expand_path('~/Library/Application Support/Sublime Text 2')
    source_path = File.join(ENV['HOME'], '.sublime_text2', 'Preferences.sublime-settings')
    target_path = File.join(root_path, 'Packages/User/Preferences.sublime-settings')

    # Make sure we have sublime folder under user home
    if !File.exists?(source_path)
      system %[cp -a #{DOTFILES_ROOT}/sublime_text2 #{ENV['HOME']}/.sublime_text2]
    end

    # Unlink current settings
    if File.exists?(target_path)
      system %[unlink \"#{target_path}\"]
    end

    # Link new settings
    Dir.chdir File.dirname(__FILE__) do
      system %[ln -vsf #{source_path} \"#{target_path}\"]
    end

    # Install Soda theme
    Dir.chdir(File.join(root_path, 'Packages')) do
      if !File.exists?("Theme - Soda")
        system('git clone https://github.com/buymeasoda/soda-theme/ "Theme - Soda"')
      else
        puts "Soda theme is already installed. Skipping."
      end
    end
  end

  desc 'Install custom themes'
  task :themes do
    repo   = 'git://github.com/daylerees/colour-schemes.git'
    path   = File.join(ENV['HOME'], '.sublime-themes')
    target = File.expand_path('~/Library/Application Support/Sublime Text 2/Packages/Custom Themes')

    # Clone or update themes repository
    system %[git clone #{repo} #{path}] if !File.exists?(path)
    system %[cd #{path} && git pull]

    # Link to Sublime home dir
    system %[unlink \"#{target}\"] if File.exists?(target)
    system %[ln -vsf #{path} \"#{target}\"]
  end
end

namespace :tools do
  desc 'Install Hub for Github'
  task :github do
    if `which hub`.to_s.empty?
      puts "Installing Hub..."
      exec("gem install hub")
    else
      puts "Hub is already installed"
    end
  end
end

desc 'Install default configuration'
task :install do
  Rake::Task['dotfiles'].invoke
  Rake::Task['binfiles'].invoke
end