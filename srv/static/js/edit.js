function Edit_BB(FileName, ID, date, content) {
  var a = document.getElementById("Edit_BB");
  document.getElementById("edit_FileName").value = FileName;
  document.getElementById("edit_ID").value = ID;
  document.getElementById("edit_date").value = date;
  document.getElementById("edit_content").value = content;
  a.style.display = "block";
}

function Insert_BB(FileName) {
  var a = document.getElementById("Insert_BB");
  document.getElementById("insert_FileName").value = FileName;
  a.style.display = "block";
}

function Edit_Personal(
  FileName,
  ID,
  date,
  content,
  partner,
  type,
  location,
  typeView,
  locationView
) {
  // alert("Mass:"+FileName+ID+date+content+partner+type+location+typeView+locationView)
  var a = document.getElementById("Edit-Personal-JournalPapers");
  document.getElementById("Personal_FileName").value = FileName;
  document.getElementById("Personal_ID").value = ID;
  document.getElementById("Personal_date").value = date;
  document.getElementById("Personal_content").value = content;
  document.getElementById("Personal_type").value = type;
  document.getElementById("Personal_location").value = location;
  document.getElementById("Personal_partner").value = partner;
  // <!--=========================================================-->
  // document.getElementById('h4-type').style.display=typeView;
  // document.getElementById('h4-location').style.display=locationView;
  // <!--=========================================================-->
  // document.getElementById('edit-type').style.display=typeView;
  // document.getElementById('edit_location').style.display=locationView;
  a.style.display = "block";
}

function Insert_Personal(FileName) {
  var a = document.getElementById("Insert_BB");
  document.getElementById("insert_FileName").value = FileName;
  a.style.display = "block";
}
