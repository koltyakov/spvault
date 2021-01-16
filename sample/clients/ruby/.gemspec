# -*- ruby -*-
# encoding: utf-8

Gem::Specification.new do |s|
  s.name          = 'spvault-client'
  s.version       = '1.0.0'
  s.authors       = ['Andrew Koltyakov']
  s.homepage      = 'https://github.com/koltyakov/spvault'
  s.summary       = 'SPVault Ruby Sample'
  s.description   = 'SPVault Ruby Sample Client'

  s.files         = `git ls-files -- ruby/*`.split("\n")
  s.executables   = `git ls-files -- ruby/main.rb`.split("\n").map do |f|
    File.basename(f)
  end
  s.require_paths = ['lib']
  s.platform      = Gem::Platform::RUBY

  s.add_dependency 'grpc', '~> 1.0'
  s.add_dependency 'date', '~> 3.1'
  s.add_dependency 'warning', '~> 1.1'
  s.add_development_dependency 'bundler', '>= 1.9'
end