{
  "Attributes": {
    "EffectiveDeliveryPolicy": "{\n\t\t\"http\": {\n\t\t\t\"defaultHealthyRetryPolicy\": {\n\t\t\t\"minDelayTarget\": 20,\n\t\t\t\"maxDelayTarget\": 20,\n\t\t\t\"numRetries\": 3,\n\t\t\t\"numMaxDelayRetries\": 0,\n\t\t\t\"numNoDelayRetries\": 0,\n\t\t\t\"numMinDelayRetries\": 0,\n\t\t\t\"backoffFunction\": \"linear\"\n\t\t\t},\n\t\t\t\"disableSubscriptionOverrides\": false\n\t\t}\n\t}",
    "Owner": "000000000000",
    "Policy": "{\n\t\t\"Version\": \"2008-10-17\",\n\t\t\"Id\": \"__default_policy_ID\",\n\t\t\"Statement\": [\n\t\t\t{\n\t\t\t\"Sid\": \"__default_statement_ID\",\n\t\t\t\"Effect\": \"Allow\",\n\t\t\t\"Principal\": {\n\t\t\t\t\"AWS\": \"*\"\n\t\t\t},\n\t\t\t\"Action\": [\n\t\t\t\t\"SNS:GetTopicAttributes\",\n\t\t\t\t\"SNS:SetTopicAttributes\",\n\t\t\t\t\"SNS:AddPermission\",\n\t\t\t\t\"SNS:RemovePermission\",\n\t\t\t\t\"SNS:DeleteTopic\",\n\t\t\t\t\"SNS:Subscribe\",\n\t\t\t\t\"SNS:ListSubscriptionsByTopic\",\n\t\t\t\t\"SNS:Publish\",\n\t\t\t\t\"SNS:Receive\"\n\t\t\t],\n\t\t\t\"Resource\": \"*\"\n\t\t\t}\n\t\t]\n\t}",
    "SubscriptionsConfirmed": "1",
    "SubscriptionsDeleted": "0",
    "SubscriptionsPending": "0",
    "TopicArn": "arn:aws:sns:us-east-1:000000000000:test-get-topic-attributes"
  },
  "ResultMetadata": {}
}