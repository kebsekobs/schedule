import { CellHelper } from "./CellHelpers";
import { createColumnHelper } from "@tanstack/react-table";
import { EditCell } from "./EditCell";

const { accessor, group } = createColumnHelper();

export const columns = [
  group({
    id: "@groupId",
    header: "Номер группы",
    columns: [
      accessor("groupId", {
        header: "",
        size: 400,
        cell: (data) => (
          <div style={{ textAlign: "center" }}>{data.getValue()}</div>
        ),
      }),
    ],
  }),
  group({
    id: "@name",
    header: "Код группы",
    columns: [
      accessor("name", {
        header: "",
        size: 50,
        cell: (data) => (
          <div style={{ textAlign: "center" }}>{data.getValue()}</div>
        ),
      }),
    ],
  }),
  group({
    id: "@GroupGrade",
    header: "Курс",
    columns: [
      accessor("name", {
        header: "",
        size: 50,
        cell: (data) => (
          <div style={{ textAlign: "center" }}>
            {CellHelper.renderGroupName(data.row.original)}
          </div>
        ),
      }),
    ],
  }),
  group({
    id: "@capacity",
    header: "Количество студентов",
    columns: [
      accessor("capacity", {
        header: "",
        size: 50,
        cell: (data) => (
          <div style={{ textAlign: "center" }}>{data.getValue()}</div>
        ),
      }),
    ],
  }),
  group({
    id: "@capacity",
    header: "",
    columns: [
      accessor("capacity", {
        header: "",
        size: 50,
        cell: (data) => <EditCell props={data.row} />,
      }),
    ],
  }),
];
