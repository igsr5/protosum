# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: proto/feature/feature_api.proto for package 'Protosum.Feature.FeatureApiPb'

require 'grpc'
require 'proto/feature/feature_api_pb'

module Protosum
  module Feature
    module FeatureApiPb
      module FeatureService
        class Service

          include ::GRPC::GenericService

          self.marshal_class_method = :encode
          self.unmarshal_class_method = :decode
          self.service_name = 'feature.feature_api_pb.FeatureService'

          rpc :GetFeature, ::Protosum::Feature::FeatureApiPb::Point, ::Protosum::Feature::FeatureApiPb::Feature
        end

        Stub = Service.rpc_stub_class
      end
    end
  end
end
