# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: post.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("post.proto", :syntax => :proto3) do
    add_message "postpb.Post" do
      optional :name, :string, 1
      optional :type, :int64, 2
      optional :user_name, :int64, 3
      optional :title, :string, 4
    end
  end
end

module Protosum
  module PostPb
    Post = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("postpb.Post").msgclass
  end
end
