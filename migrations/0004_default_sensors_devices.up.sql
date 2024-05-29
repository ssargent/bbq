update bbq.sensors 
set is_default = true
where id = uuid_nil();

insert into bbq.devices
(id, name, location, is_default)
values
(uuid_nil(), 'default-device', 'Back Deck', true);