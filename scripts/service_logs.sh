oklog stream | sed -u 's|.*{|{|' | grep --line-buffered '^{' | jq '{ "service": .service, "out": . }'
