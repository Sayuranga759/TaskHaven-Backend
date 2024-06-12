--
-- ***START ENUM TYPE***
-- At the very first time, after the DB creation need to run bellow create ENUM script at once.
--
-- Create type ***status***
CREATE TYPE status AS ENUM ('completed', 'to_do', 'on_hold');
--
-- Create type  ***priority_level***
CREATE TYPE priority_level AS ENUM ('low', 'medium', 'high');
--
-- ***END ENUM TYPE***
--