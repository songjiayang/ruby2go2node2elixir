require 'sinatra/base'
require "sinatra/json"

class User
  include Mongoid::Document
  field :username, type: String
  field :created_at, type: DateTime
  field :updated_at, type: DateTime
end


class App < Sinatra::Base
  set :sessions, false
  set :environment, :production

  get "/users" do
    json User.all
  end

  get "/users/:id" do
    json User.find(params['id'])
  end

  get "/static" do
    data = {
  		"_id":         "57472be59cf6a4c96b088a17",
  		"created_at": "2016-05-26T20:08:13.494+08:00",
  		"updated_at": "2016-05-26T20:08:13.494+08:00",
  		"username":   "josh whilist",
    }
    json data
  end
end
