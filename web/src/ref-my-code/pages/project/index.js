import ItemList from "../../component/ItemList";
export default function ProjectPage(props) {
  const url = "https://api.github.com/repos/asiashaowei/PublicNote/contents/";
  console.log("==>>");
  console.log(props.breadcrumbItems);
  return (
    <div>
      <ItemList
        breadcrumbItems={props.breadcrumbItems}
        setBreadcrumbItems={props.setBreadcrumbItems}
        url={url}
      />
    </div>
  );
}
// 進度 無法更改導覽列內容
