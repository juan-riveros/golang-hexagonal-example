drop table if exists KeyValues;

create table KeyValues (
    id int primary key,
    key text,
    value text,
    unique(key, value)
);

insert into KeyValues (key, value) values ("agegroup", "young");
insert into KeyValues (key, value) values ("agegroup", "adult");
insert into KeyValues (key, value) values ("agegroup", "midlife");
insert into KeyValues (key, value) values ("agegroup", "senior");
insert into KeyValues (key, value) values ("agegroup", "elderly");
insert into KeyValues (key, value) values ("agegroup", "ancient");
insert into KeyValues (key, value) values ("agegroup", "centenarian");
insert into KeyValues (key, value) values ("agegroup", "living dead");
insert into KeyValues (key, value) values ("agegroup", "ghost");

insert into KeyValues (key, value) values ("gender", "male");
insert into KeyValues (key, value) values ("gender", "female");

insert into KeyValues (key, value) values ("locale", "beach");
insert into KeyValues (key, value) values ("locale", "mountain");
insert into KeyValues (key, value) values ("locale", "side of the road");
insert into KeyValues (key, value) values ("locale", "city hall");
insert into KeyValues (key, value) values ("locale", "school");
insert into KeyValues (key, value) values ("locale", "office building");

insert into KeyValues (key, value) values ("mood", "happy");
insert into KeyValues (key, value) values ("mood", "sad");
insert into KeyValues (key, value) values ("mood", "angry");
insert into KeyValues (key, value) values ("mood", "somber");
insert into KeyValues (key, value) values ("mood", "hectic");
insert into KeyValues (key, value) values ("mood", "quiet");
insert into KeyValues (key, value) values ("mood", "tense");

insert into KeyValues (key, value) values ("profession", "teacher");
insert into KeyValues (key, value) values ("profession", "officer");
insert into KeyValues (key, value) values ("profession", "maid");
insert into KeyValues (key, value) values ("profession", "official");
insert into KeyValues (key, value) values ("profession", "ranger");
insert into KeyValues (key, value) values ("profession", "sales associate");

insert into KeyValues (key, value) values ("timeofday", "morning");
insert into KeyValues (key, value) values ("timeofday", "noon");
insert into KeyValues (key, value) values ("timeofday", "afternoon");
insert into KeyValues (key, value) values ("timeofday", "night");

insert into KeyValues (key, value) values ("timeperiod", "1930s");
insert into KeyValues (key, value) values ("timeperiod", "prehistoric");
insert into KeyValues (key, value) values ("timeperiod", "modern");
insert into KeyValues (key, value) values ("timeperiod", "20XX");
insert into KeyValues (key, value) values ("timeperiod", "turn of the 20th century");


