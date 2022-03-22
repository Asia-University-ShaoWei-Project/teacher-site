<div id="PDF" class="modal" style="z-index: 99">
  <form class="modal-content-pdf animate">
    <div style="height: 90%;position: relative;">
      <!--            <iframe src=""></iframe>-->
      <iframe id="iframe-pdf"
        style="height: 100%;width: 90%;border:1px solid DodgerBlue;margin: 2% 3% auto 3%"></iframe>
      <span style="position: absolute;right: 2.5%;top: 30px"
        onclick="document.getElementById('PDF').style.display='none'" class="close" title="Close Modal">X</span>
      <span id="PDF_Button" style="line-height:100px;position: absolute;margin: 80px auto">
        <a href="#" onclick="document.getElementById('PDF')
                .style.display='none'"><i class="fa fa-download" title="Download"></i></a>
        <br>
        <a href="#" onclick="document.getElementById('PDF')
                .style.display='none'"><i class="fa fa-expand" title="Amplification"></i></a>
        <br>
      </span>
    </div>
  </form>
</div>