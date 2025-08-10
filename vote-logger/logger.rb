require 'sinatra'
require 'json'

post '/log' do
  data = JSON.parse(request.body.read)
  File.open('votes.log', 'a') do |f|
    f.puts "#{data['time']} - Voted: #{data['option']}"
  end
  status 200
end