import { createColumnHelper } from "@tanstack/react-table";
import { EditCell } from "./EditCell";

const { accessor, group } = createColumnHelper();

export const columns = [
  group({
    id: "@id",
    header: "Номер аудитории",
    columns: [
      accessor("classroomId", {
        header: "",
        size: 400,
        cell: (data) => (
          <div style={{ textAlign: "center" }}>{data.getValue()}</div>
        ),
      }),
    ],
  }),
  group({
    id: "@capacity",
    header: "Вместимость аудитории",
    columns: [
      accessor("capacity", {
        header: "",
        size: 100,
        cell: (data) => (
          <div style={{ textAlign: "center" }}>{data.getValue()}</div>
        ),
      }),
    ],
  }),
  group({
    id: "@edit",
    header: "",
    columns: [
      accessor("edit", {
        header: "",
        size: 50,
        cell: (data) => <EditCell props={data.row} />,
      }),
    ],
  }),
];
