# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/feature/feature_api.proto

require 'google/protobuf'

require 'google/api/annotations_pb'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("proto/feature/feature_api.proto", :syntax => :proto3) do
    add_message "feature.feature_api_pb.Point" do
      optional :latitude, :int32, 1
    end
    add_message "feature.feature_api_pb.Feature" do
      optional :code, :int32, 1
    end
  end
end

module Protosum
  module Feature
    module FeatureApiPb
      Point = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("feature.feature_api_pb.Point").msgclass
      Feature = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("feature.feature_api_pb.Feature").msgclass
    end
  end
end
