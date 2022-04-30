const infoBoxElem = document.getElementById("info-box");
const coursesBoxElem = document.getElementById("courses-box");
const optionIcons = {
  INFO: `<i class="fa fa-calendar" aria-hidden="true"></i>`,
  COURSE: `<i class="fa fa-hand-o-right" aria-hidden="true"></i>`,
};
function showOptionButtons(data) {
  console.log(`show option buttons. data length: ${data.length}`);
  let infoIndex = 0;
  let infoOptionElem = createOptionButton(infoIndex, optionIcons.INFO);
  infoBoxElem.appendChild(infoOptionElem);
  for (let i = 1; i < data.length; i++) {
    elem = createOptionButton(i, optionIcons.COURSE);
    coursesBoxElem.appendChild(elem);
  }
}
function createOptionButton(index, iconElem) {
  let optionItemElem = document.createElement("li");
  let optionBtnElem = document.createElement("a");
  let optionTitleElem = document.createElement("span");
  let title = items[index].getNameZh() + brTag + items[index].getNameUs();
  optionItemElem.className = `nav-item`;
  optionBtnElem.className = `nav-link`;
  optionTitleElem.className = `option-title`;
  optionBtnElem.href = "#";

  optionTitleElem.innerHTML = iconElem + " " + title;
  optionBtnElem.onclick = () => items[index].updateData();
  optionBtnElem.appendChild(optionTitleElem);
  optionItemElem.appendChild(optionBtnElem);
  return optionItemElem;
}
