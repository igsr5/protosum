# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: feature/feature_api.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("feature/feature_api.proto", :syntax => :proto3) do
    add_message "feature.feature_api_pb.Point" do
      optional :latitude, :int32, 1
      optional :longitude, :int32, 2
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
