import { createColumnHelper } from "@tanstack/react-table";
import { EditCell } from "./EditCell";

const { accessor, group } = createColumnHelper();

export const columns = [
  group({
    id: "@name",
    header: "Преподаватель",
    columns: [
      accessor("name", {
        header: "",
        size: 400,
        cell: (data) => (
          <div style={{ textAlign: "center" }}>{data.getValue()}</div>
        ),
      }),
    ],
  }),
  group({
    id: "@editCell",
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
