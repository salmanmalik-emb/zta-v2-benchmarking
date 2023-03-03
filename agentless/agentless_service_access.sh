#!/bin/sh

cookie="Cookie:sessionid=o709l22u91myk2tcnsox62tkgdfnvgwb;access-token=52da7f32-3990-4e2a-a320-030a0bf36164;service_id=d2520937-57e1-420e-a195-66bd9d1cbae1"

echo "Cookies: $cookie" 

curl -s -L -v -O "https://c929bc8ae5fdd977fc05545be7eeafb7-agentless.feature-4.extremecloudztna.com/images/bigtest.jpg" -H "$cookie" -H "User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.96 Safari/537.36" -H "X-FORWARDED-FOR: 13.250.241.161" --limit-rate 80K