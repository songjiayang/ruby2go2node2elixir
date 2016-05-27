$LOAD_PATH.unshift(".") unless $LOAD_PATH.include?(".")

require "mongoid"
Mongoid.load!("mongoid.yml", :production)

require 'app'
run App
