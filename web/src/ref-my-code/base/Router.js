// icons
import PersonIcon from "@material-ui/icons/Person";
import CreateIcon from "@material-ui/icons/Create";
import DescriptionIcon from "@material-ui/icons/Description";
import SchoolIcon from "@material-ui/icons/School";
// pages
import IntrPage from "../pages/intr/index";
import NotePage from "../pages/note/index";
import ProjectPage from "../pages/project/index";
import UnivPage from "../pages/univ/index";

export const router_config = [
  {
    path: "/intr",
    title: "Personal Profile",
    icon: <PersonIcon fontSize={"small"} />,
    page: (breadcrumbItems, setBreadcrumbItems) => <IntrPage />,
  },
  {
    path: "/note",
    title: "Note",
    icon: <CreateIcon fontSize={"small"} />,
    page: (breadcrumbItems, setBreadcrumbItems) => <NotePage />,
  },
  {
    path: "/project",
    title: "Project",
    icon: <DescriptionIcon fontSize={"small"} />,
    page: (breadcrumbItems, setBreadcrumbItems) => (
      <ProjectPage
        breadcrumbItems={breadcrumbItems}
        setBreadcrumbItems={setBreadcrumbItems}
      />
    ),
  },
  {
    path: "/univ",
    title: "Asia University",
    icon: <SchoolIcon fontSize={"small"} />,
    page: (breadcrumbItems, setBreadcrumbItems) => <UnivPage />,
  },
];
