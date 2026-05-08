{
  "Attributes": {
    "EffectiveDeliveryPolicy": "{\n\t\t\"http\": {\n\t\t\t\"defaultHealthyRetryPolicy\": {\n\t\t\t\"minDelayTarget\": 20,\n\t\t\t\"maxDelayTarget\": 20,\n\t\t\t\"numRetries\": 3,\n\t\t\t\"numMaxDelayRetries\": 0,\n\t\t\t\"numNoDelayRetries\": 0,\n\t\t\t\"numMinDelayRetries\": 0,\n\t\t\t\"backoffFunction\": \"linear\"\n\t\t\t},\n\t\t\t\"disableSubscriptionOverrides\": false\n\t\t}\n\t}",
    "Owner": "000000000000",
    "Policy": "{\"Version\":\"2008-10-17\",\"ID\":\"us-east-1/000000000000/test-get-topic-attributes\",\"Statement\":[{\"Sid\":\"us-east-1/000000000000/test-get-topic-attributes\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"*\"},\"Action\":[\"SNS:GetTopicAttributes\",\"SNS:SetTopicAttributes\",\"SNS:AddPermission\",\"SNS:RemovePermission\",\"SNS:DeleteTopic\",\"SNS:Subscribe\",\"SNS:ListSubscriptionsByTopic\",\"SNS:Publish\"],\"Resource\":\"arn:aws:sns:us-east-1:000000000000:test-get-topic-attributes\",\"Condition\":{\"StringLike\":{\"AWS:SourceArn\":\"arn:aws:*:*:000000000000:*\"}}}]}",
    "SubscriptionsConfirmed": "1",
    "SubscriptionsDeleted": "0",
    "SubscriptionsPending": "0",
    "TopicArn": "arn:aws:sns:us-east-1:000000000000:test-get-topic-attributes"
  },
  "ResultMetadata": {}
}