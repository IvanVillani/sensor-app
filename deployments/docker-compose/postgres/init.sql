create table device (
  id char(5) PRIMARY KEY NOT NULL,
  name varchar(20) NOT NULL,
  description text NOT NULL
);

create type sensorgroups as enum ('CPU_TEMP', 'CPU_USAGE', 'MEMORY_USAGE');

create table sensor (
  id char(4) PRIMARY KEY NOT NULL,
  name varchar(20) NOT NULL,
  description text NOT NULL,
  unit varchar(20) NOT NULL,
  sensor_groups sensorgroups NOT NULL,
  device_id char(5) NOT NULL,
  CONSTRAINT fk_device
      FOREIGN KEY(device_id) 
	    REFERENCES device(id)
);

insert into device (id, name, description) values
('D1111', 'MBP13', 'MacBook Pro 13'),
('D2222', 'MBP16', 'MacBook Pro 16'),
('D3333', 'MBA2019', 'MacBook Pro 2019'),
('MBP13', 'MBA20121', 'MacBook Pro 2019 - WORK');

insert into sensor (id, name, description, unit, sensor_groups, device_id) values
('S111', 'TCT0', 'CPU temperature sensor', 'Celsius', 'CPU_TEMP', 'D1111'),
('S112', 'TCT0', 'CPU temperature sensor', 'Celsius', 'CPU_TEMP', 'D2222'),
('S113', 'TCT0', 'CPU temperature sensor', 'Celsius', 'CPU_TEMP', 'D3333'),
('S221', 'TCU0', 'CPU usage sensor', 'Percent', 'CPU_USAGE', 'D1111'),
('S222', 'TCU0', 'CPU usage sensor', 'Percent', 'CPU_USAGE', 'D2222'),
('S223', 'TCU0', 'CPU usage sensor', 'Percent', 'CPU_USAGE', 'D3333'),
('S331', 'TMU0', 'Memory usage sensor', 'Percent', 'MEMORY_USAGE', 'D1111'),
('S332', 'TMU0', 'Memory usage sensor', 'Percent', 'MEMORY_USAGE', 'D2222'),
('S333', 'TMU0', 'Memory usage sensor', 'Percent', 'MEMORY_USAGE', 'D3333'),
('TC0P', 'TC0P', 'CPU temperature sensor', 'Celsius', 'CPU_TEMP', 'MBP13');
