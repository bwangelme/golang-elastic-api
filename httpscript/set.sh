#!/bin/bash
#
# Author: bwangel<bwangel.me@gmail.com>
# Date: 5,10,2020 13:33

source ./bash.sh

http --json post "$BASE_URL/set" \
 id="1" \
 parent_id="" \
 operation="add" \
 source='{"cn_title": "\u676f\u9152\u4eba\u751f", "year": "2004", "id": "1291833", "is_comment_hidden": false, "ratings": [131, 867, 10656, 22773, 10707], "genres": ["\u5267\u60c5", "\u559c\u5267", "\u7231\u60c5"], "is_review_hidden": false, "countries": ["\u7f8e\u56fd", "\u5308\u7259\u5229"], "casts": ["1017893", "1019041", "1049512", "1010565", "1386061", "1009416", "1386057", "1100584", "1009582", "1320238", "1135460", "1074663", "1031954", "1085782", "1122417", "1083839", "1345759", "1205699", "1286995"], "release_date": "2004-09-13", "picture_url": "http://img1.doubanio.com/view/photo/s_ratio_poster/public/p1910902469.jpg", "mainland_movie_length": "126\u5206\u949f", "languages": ["\u82f1\u8bed", "\u4e9a\u7f8e\u5c3c\u4e9a\u8bed"], "directors": ["1004720"], "score": "7.9", "is_rated_enough": true, "pubdates": ["2004-09-13(\u591a\u4f26\u591a\u7535\u5f71\u8282)", "2005-01-21(\u7f8e\u56fd)", "2005-02-03(\u5308\u7259\u5229)"], "title": "\u676f\u9152\u4eba\u751f", "is_hidden_for_anonymous": false, "is_tv": false}'
