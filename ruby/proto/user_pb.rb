# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/user.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("proto/user.proto", :syntax => :proto3) do
    add_message "userpb.User" do
      optional :name, :string, 1
    end
  end
end

module UserPb
  User = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("userpb.User").msgclass
end
