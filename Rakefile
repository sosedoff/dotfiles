require 'yaml'
require 'fileutils'
require 'pp'

DOTFILES_HOME = ENV['HOME']
DOTFILES_ROOT = File.dirname(__FILE__)
CONFIG_PATH   = File.join(DOTFILES_ROOT, 'config.yml')
CONFIG        = YAML.load_file(CONFIG_PATH)

desc "Show current config"
task :config do
  CONFIG.each_pair do |name, files|
    puts "#{name}:"
    files.each { |f| puts " - #{f}" }
  end
end

desc "Install dot files"
task :dotfiles do
  files     = CONFIG['dotfiles']
  files_dir = File.join(DOTFILES_ROOT, 'dotfiles')

  files.each do |name|
    path      = File.join(DOTFILES_ROOT, 'dotfiles', name)
    link_path = File.join(DOTFILES_HOME, ".#{name}")

    system "unlink #{link_path}" if File.exists?(link_path)
    system "ln -vsf #{path} #{link_path}"
  end
end

desc "Install bin files"
task :binfiles do
  files     = CONFIG['binfiles']
  files_dir = File.join(DOTFILES_ROOT, 'bin')

  FileUtils.mkdir_p("#{DOTFILES_HOME}/bin")

  files.each do |name|
    path      = File.join(DOTFILES_ROOT, 'bin', name)
    link_path = "#{DOTFILES_HOME}/bin/#{File.basename(name)}"

    system "unlink #{link_path}" if File.exists?(link_path)
    system "ln -vsf #{path} #{link_path}"
  end
end

namespace :sublime do
  desc "Install ST2 settings"
  task :settings do
    root_path   = File.expand_path('~/Library/Application Support/Sublime Text 3')
    source_path = File.join(ENV['HOME'], '.sublime_text', 'Preferences.sublime-settings')
    target_path = File.join(root_path, 'Packages/User/Preferences.sublime-settings')

    # Make sure we have sublime folder under user home
    if !File.exists?(source_path)
      system %[cp -a #{DOTFILES_ROOT}/sublime_text3 #{ENV['HOME']}/.sublime_text]
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

  desc "Install ST2 themes"
  task :themes do
    repo   = "git://github.com/daylerees/colour-schemes.git"
    path   = File.join(ENV["HOME"], ".sublime-themes")
    target = File.expand_path("~/Library/Application Support/Sublime Text 3/Packages/Custom Themes")

    # Clone or update themes repository
    system %[git clone #{repo} #{path}] if !File.exists?(path)
    system %[cd #{path} && git pull]

    # Link to Sublime home dir
    system %[unlink \"#{target}\"] if File.exists?(target)
    system %[ln -vsf #{path} \"#{target}\"]
  end
end

task :gvm do
  exec("bash < <(curl -s https://raw.github.com/moovweb/gvm/master/binscripts/gvm-installer)")
end

desc "Install all settings"
task :install do
  Rake::Task['dotfiles'].invoke
  Rake::Task['binfiles'].invoke
end