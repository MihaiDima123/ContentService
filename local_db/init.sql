create table tenant (
    id SERIAL primary key,
    tenant_ref varchar(50) unique,
    channel varchar(20),
    language varchar(5),
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    updated_at DATE
);

create table configuration (
    id SERIAL primary key,
    Configuration_Name varchar(30) unique,
    Configuration_Description varchar(50)
);

create table tenant_configuration (
    id SERIAL primary key,
    tenant_id int,
    configuration_id int,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    updated_at DATE,
    constraint fk_tenant_configuration_tenant foreign key (tenant_id) references tenant(id),
    constraint fk_tenant_configuration_configuration foreign key (configuration_id) references configuration(id)
);

insert into tenant (tenant_ref, channel, language) values ('blogs-ro_RO', 'blogs', 'ro_RO');