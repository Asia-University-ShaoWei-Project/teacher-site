// window.addEventListener()
function test(FileName, Num, Chose) {
  var pdf_elem = document.getElementById("PDF");
  var pdf_iframe_elem = document.getElementById("iframe-pdf");
  try {
    //0是上 1下
    switch (Num) {
      case 0: {
        pdf_iframe_elem.src = "PcWeb/CourseSlides/" + Chose + "/" + FileName;
        break;
      }
      case 1: {
        pdf_iframe_elem.src = "PcWeb/Homeworks/" + Chose + "/Hw" + FileName;
        break;
      }
    }
  } catch (e) {
    alert("?" + e.getErrorMessage());
  }
  pdf_elem.style.display = "block";
}

$(".side-bar-option").click(function () {
  $(".index-news-icon").fadeIn(1400);
  $(".index-news-list-info-box").fadeIn(1950);
  // $("#CENTER").fadeIn(1950);
});

$(document).ready(function () {
  const sign_in_frame_elem = document.getElementById("sign-in-frame");
  $("#sign-in-exit-button").click(function () {
    sign_in_frame_elem.fadeOut(1000);
  });
  $("#sign-in-header-button").click(function () {
    $("#sign-in-frame").fadeIn(1200);
  });
});

window.onresize = function () {
  WebSize();
};

function WebSize() {
  try {
    var Width = $(window).width();
    if (Width >= 1049) {
      if (Width >= 1501) {
        ReSet_size(Width * 0.26, Width * 0.74);
      } else {
        ReSet_size(Width * 0.3, Width * 0.7);
      }
    } else {
      ReSet_size(0, Width);
    }
  } catch (e) {
    alert("Error" + e.getErrorMessage());
  }
}

function ReSet_size(LEFT_W, RIGHT_W) {
  var DIV_Left = document.getElementById("LEFT");
  DIV_Left.style.width = LEFT_W + "px";
  var DIV_Center = document.getElementById("CENTER");
  // var DIV_Center1 = document.getElementById("content");
  DIV_Center.style.width = RIGHT_W + "px";
  DIV_Center.style.maxWidth = RIGHT_W + "px";
  // DIV_Center1.style.width=RIGHT_W+"px";
  // DIV_Center1.style.maxWidth=RIGHT_W+"px";
  // DIV_Center.style.maxWidth=center_width+"px";
  DIV_Center.style.left = DIV_Left.style.width;
  // alert(DIV_Center.style.left);
}
//滾輪移動
// function scrollFunction() {
//     if(document.body.scrollTop>20 || document.documentElement.scrollTop>20){
//         document.getElementById("header").style.height="80px";
//     }else {
//         document.getElementById("header").style.height="65px";
//     }
// }
//
// -----------------
// var btn = document.createElement("button");
// btn.innerHTML = "something";
// obj.appendChild()
