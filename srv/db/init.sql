DROP TABLE IF EXISTS auths;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS teachers;
CREATE TABLE auths(
  id VARCHAR(50) NOT NULL,
  password TEXT NOT NULL,
  PRIMARY KEY(id)
);
CREATE TABLE tokens(
  auth_id VARCHAR(50),
  teacher_id INTEGER,
  token VARCHAR(255) NOT NULL,
  FOREIGN KEY (auth_id) REFERENCES auths(id),
  FOREIGN KEY (teacher_id) REFERENCES teachers(id)
);
CREATE TABLE teachers (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(50) NOT NULL,
  -- Email
  email VARCHAR(50) NOT NULL,
  -- 辦公室
  office VARCHAR(10) NOT NULL,
  -- 分機
  call VARCHAR(20) NOT NULL,
  -- 學歷
  education VARCHAR(50) NOT NULL
  -- create_date DATE,
  -- update_date DATE,
);
-- TABLE courses (
--   course_id INT,
--   content_id INT,
--   title VARCHAR(50) NOT NULL,
--   FOREIGN KEY(course_id) REFERENCES tearchs(id)
--   FOREIGN KEY(content_id) REFERENCES content(id)
-- );
-- "MessageAnnouncement", "PersonalInformation", "ComputerNetworks", "ComputerOrganization"


INSERT INTO auths(id, password)VALUES(
  'rikki',
  '$2y$08$G4qPhq8U2bftPnsBSKeOBOGby2zQL/vP/y3sFPYILr8dH0DMGt.IS'
);
INSERT INTO tokens(auth_id, teacher_id, token)VALUES(
  'rikki',
  1,
  '03-05-246235f13480594'
);
INSERT INTO teachers(name, email, office, call, education)
VALUES(
  'Jui-Chi Chen (陳瑞奇)(Rikki)',
  'rikki@asia.edu.tw',
  'HB13',
  '(04)2332-3456 ext. 20013',
  '國立中興大學資訊科學博士'
);
SELECT * FROM auths;
SELECT * FROM tokens;
SELECT * FROM teachers;