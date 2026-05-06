package sns

import (
	"context"
	"os"
	"testing"
)

const baseURL = "http://localhost:4566"

func TestGetTopicAttributes(t *testing.T) {
	var opts []Option
	if dir := os.Getenv("KUMO_DATA_DIR"); dir != "" {
		opts = append(opts, WithDataDir(dir))
	}

	storage := NewMemoryStorage(baseURL, opts...) // Or however you initialize your storage
	ctx := context.Background()
	topicName := "TestTopic"
	inputAttributes := map[string]string{"DisplayName": "MyTopic"}

	// Create a topic with attributes
	topic, err := storage.CreateTopic(ctx, topicName, inputAttributes)
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	// Retrieve attributes
	got, err := storage.GetTopicAttributes(ctx, topic.ARN)

	// Assertions
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(inputAttributes) {
		t.Errorf("expected %d attributes, got %d", len(inputAttributes), len(got))
	}

	for k, expectedV := range inputAttributes {
		if gotV, exists := got[k]; !exists || gotV != expectedV {
			t.Errorf("attribute %s: expected %s, got %s", k, expectedV, gotV)
		}
	}

	// Verify encapsulation
	// Modify the returned map and ensure storage is not affected
	got["DisplayName"] = "HackedValue"

	verifyStorage, _ := storage.GetTopicAttributes(ctx, topic.ARN)
	if verifyStorage["DisplayName"] == "HackedValue" {
		t.Error("Data Leak detected: modifying returned attributes modified internal storage!")
	}
}

func TestSetTopicAttributes(t *testing.T) {
	var opts []Option
	if dir := os.Getenv("KUMO_DATA_DIR"); dir != "" {
		opts = append(opts, WithDataDir(dir))
	}

	storage := NewMemoryStorage(baseURL, opts...) // Or however you initialize your storage
	ctx := context.Background()
	topicName := "TestTopic"
	emptyAttribute := make(map[string]string)
	inputAttributeName := "DisplayName"
	inputAttributeValue := "test"

	// Create a topic with attributes
	topic, err := storage.CreateTopic(ctx, topicName, emptyAttribute)
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}

	err = storage.SetTopicAttributes(ctx, topic.ARN, inputAttributeName, inputAttributeValue)
	if err != nil {
		t.Fatalf("Failed to set topic attributes, got %v", err)
	}

	// Retrieve attributes
	got, err := storage.GetTopicAttributes(ctx, topic.ARN)

	// Assertions
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != 1 {
		t.Errorf("expected %s attributes, got %d", "1", len(got))
	}

	if gotV, exists := got[inputAttributeName]; !exists || gotV != inputAttributeValue {
		t.Errorf("attribute %s: expected %s, got %s", inputAttributeName, inputAttributeValue, gotV)
	}
}
