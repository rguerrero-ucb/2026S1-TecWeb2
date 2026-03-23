create table todo.tareas (
    id serial primary key,
    nombre varchar(255) not null,
    estado smallint not null
);

insert into todo.tareas(nombre, estado) values
('tarea1', 0),
('tarea2', 0);

select * from todo.tareas;