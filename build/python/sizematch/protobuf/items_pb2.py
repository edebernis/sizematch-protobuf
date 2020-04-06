# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: items.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='items.proto',
  package='sizematch.protobuf.items',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\x0bitems.proto\x12\x18sizematch.protobuf.items\"\xa7\x02\n\x04Item\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0e\n\x06source\x18\x02 \x01(\t\x12\x0c\n\x04lang\x18\x03 \x01(\t\x12\x0c\n\x04urls\x18\x04 \x03(\t\x12\x0c\n\x04name\x18\x05 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12\x12\n\ncategories\x18\x07 \x03(\t\x12\x12\n\nimage_urls\x18\x08 \x03(\t\x12\x42\n\ndimensions\x18\t \x03(\x0b\x32..sizematch.protobuf.items.Item.DimensionsEntry\x12\r\n\x05price\x18\n \x01(\x04\x12\x16\n\x0eprice_currency\x18\x0b \x01(\t\x1a\x31\n\x0f\x44imensionsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xc4\x02\n\x0eNormalizedItem\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0e\n\x06source\x18\x02 \x01(\t\x12;\n\x04lang\x18\x03 \x01(\x0e\x32-.sizematch.protobuf.items.NormalizedItem.Lang\x12\x0c\n\x04urls\x18\x04 \x03(\t\x12\x0c\n\x04name\x18\x05 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12\x12\n\ncategories\x18\x07 \x03(\t\x12\x12\n\nimage_urls\x18\x08 \x03(\t\x12\x38\n\ndimensions\x18\t \x01(\x0b\x32$.sizematch.protobuf.items.Dimensions\x12.\n\x05price\x18\n \x01(\x0b\x32\x1f.sizematch.protobuf.items.Price\"\x16\n\x04Lang\x12\x06\n\x02\x45N\x10\x00\x12\x06\n\x02\x46R\x10\x01\":\n\nDimensions\x12\x0e\n\x06height\x18\x01 \x01(\x04\x12\r\n\x05width\x18\x02 \x01(\x04\x12\r\n\x05\x64\x65pth\x18\x03 \x01(\x04\"p\n\x05Price\x12\r\n\x05price\x18\x01 \x01(\x04\x12:\n\x08\x63urrency\x18\x02 \x01(\x0e\x32(.sizematch.protobuf.items.Price.Currency\"\x1c\n\x08\x43urrency\x12\x07\n\x03\x45UR\x10\x00\x12\x07\n\x03GBP\x10\x01\x62\x06proto3'
)



_NORMALIZEDITEM_LANG = _descriptor.EnumDescriptor(
  name='Lang',
  full_name='sizematch.protobuf.items.NormalizedItem.Lang',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='EN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FR', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=642,
  serialized_end=664,
)
_sym_db.RegisterEnumDescriptor(_NORMALIZEDITEM_LANG)

_PRICE_CURRENCY = _descriptor.EnumDescriptor(
  name='Currency',
  full_name='sizematch.protobuf.items.Price.Currency',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='EUR', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GBP', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=810,
  serialized_end=838,
)
_sym_db.RegisterEnumDescriptor(_PRICE_CURRENCY)


_ITEM_DIMENSIONSENTRY = _descriptor.Descriptor(
  name='DimensionsEntry',
  full_name='sizematch.protobuf.items.Item.DimensionsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='sizematch.protobuf.items.Item.DimensionsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='sizematch.protobuf.items.Item.DimensionsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=288,
  serialized_end=337,
)

_ITEM = _descriptor.Descriptor(
  name='Item',
  full_name='sizematch.protobuf.items.Item',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='sizematch.protobuf.items.Item.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='sizematch.protobuf.items.Item.source', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lang', full_name='sizematch.protobuf.items.Item.lang', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='urls', full_name='sizematch.protobuf.items.Item.urls', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='sizematch.protobuf.items.Item.name', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='sizematch.protobuf.items.Item.description', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='categories', full_name='sizematch.protobuf.items.Item.categories', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image_urls', full_name='sizematch.protobuf.items.Item.image_urls', index=7,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dimensions', full_name='sizematch.protobuf.items.Item.dimensions', index=8,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='sizematch.protobuf.items.Item.price', index=9,
      number=10, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price_currency', full_name='sizematch.protobuf.items.Item.price_currency', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ITEM_DIMENSIONSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=42,
  serialized_end=337,
)


_NORMALIZEDITEM = _descriptor.Descriptor(
  name='NormalizedItem',
  full_name='sizematch.protobuf.items.NormalizedItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='sizematch.protobuf.items.NormalizedItem.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='sizematch.protobuf.items.NormalizedItem.source', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lang', full_name='sizematch.protobuf.items.NormalizedItem.lang', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='urls', full_name='sizematch.protobuf.items.NormalizedItem.urls', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='sizematch.protobuf.items.NormalizedItem.name', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='sizematch.protobuf.items.NormalizedItem.description', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='categories', full_name='sizematch.protobuf.items.NormalizedItem.categories', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image_urls', full_name='sizematch.protobuf.items.NormalizedItem.image_urls', index=7,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dimensions', full_name='sizematch.protobuf.items.NormalizedItem.dimensions', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='sizematch.protobuf.items.NormalizedItem.price', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _NORMALIZEDITEM_LANG,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=340,
  serialized_end=664,
)


_DIMENSIONS = _descriptor.Descriptor(
  name='Dimensions',
  full_name='sizematch.protobuf.items.Dimensions',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='height', full_name='sizematch.protobuf.items.Dimensions.height', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='width', full_name='sizematch.protobuf.items.Dimensions.width', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='depth', full_name='sizematch.protobuf.items.Dimensions.depth', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=666,
  serialized_end=724,
)


_PRICE = _descriptor.Descriptor(
  name='Price',
  full_name='sizematch.protobuf.items.Price',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='price', full_name='sizematch.protobuf.items.Price.price', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='currency', full_name='sizematch.protobuf.items.Price.currency', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _PRICE_CURRENCY,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=726,
  serialized_end=838,
)

_ITEM_DIMENSIONSENTRY.containing_type = _ITEM
_ITEM.fields_by_name['dimensions'].message_type = _ITEM_DIMENSIONSENTRY
_NORMALIZEDITEM.fields_by_name['lang'].enum_type = _NORMALIZEDITEM_LANG
_NORMALIZEDITEM.fields_by_name['dimensions'].message_type = _DIMENSIONS
_NORMALIZEDITEM.fields_by_name['price'].message_type = _PRICE
_NORMALIZEDITEM_LANG.containing_type = _NORMALIZEDITEM
_PRICE.fields_by_name['currency'].enum_type = _PRICE_CURRENCY
_PRICE_CURRENCY.containing_type = _PRICE
DESCRIPTOR.message_types_by_name['Item'] = _ITEM
DESCRIPTOR.message_types_by_name['NormalizedItem'] = _NORMALIZEDITEM
DESCRIPTOR.message_types_by_name['Dimensions'] = _DIMENSIONS
DESCRIPTOR.message_types_by_name['Price'] = _PRICE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Item = _reflection.GeneratedProtocolMessageType('Item', (_message.Message,), {

  'DimensionsEntry' : _reflection.GeneratedProtocolMessageType('DimensionsEntry', (_message.Message,), {
    'DESCRIPTOR' : _ITEM_DIMENSIONSENTRY,
    '__module__' : 'items_pb2'
    # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.Item.DimensionsEntry)
    })
  ,
  'DESCRIPTOR' : _ITEM,
  '__module__' : 'items_pb2'
  # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.Item)
  })
_sym_db.RegisterMessage(Item)
_sym_db.RegisterMessage(Item.DimensionsEntry)

NormalizedItem = _reflection.GeneratedProtocolMessageType('NormalizedItem', (_message.Message,), {
  'DESCRIPTOR' : _NORMALIZEDITEM,
  '__module__' : 'items_pb2'
  # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.NormalizedItem)
  })
_sym_db.RegisterMessage(NormalizedItem)

Dimensions = _reflection.GeneratedProtocolMessageType('Dimensions', (_message.Message,), {
  'DESCRIPTOR' : _DIMENSIONS,
  '__module__' : 'items_pb2'
  # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.Dimensions)
  })
_sym_db.RegisterMessage(Dimensions)

Price = _reflection.GeneratedProtocolMessageType('Price', (_message.Message,), {
  'DESCRIPTOR' : _PRICE,
  '__module__' : 'items_pb2'
  # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.Price)
  })
_sym_db.RegisterMessage(Price)


_ITEM_DIMENSIONSENTRY._options = None
# @@protoc_insertion_point(module_scope)
