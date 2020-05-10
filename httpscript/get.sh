#!/bin/bash
#
# Author: bwangel<bwangel.me@gmail.com>
# Date: 5,10,2020 13:33

source ./bash.sh

http post "$BASE_URL/get" << "END"
{
    "query_type": "",
    "child_type": "",
    "start": 0,
    "size": 10,
    "sort": {
        "field": "ratings",
        "asc": true
    },
    "query_json": [
        {
            "query_type": "must",
            "match": "text",
            "key": "title",
            "value": "酒"
        }
    ]
}
END
