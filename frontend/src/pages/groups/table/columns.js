import { CellHelper } from "./CellHelpers";
import { createColumnHelper } from "@tanstack/react-table";
import { EditCell } from "./EditCell";
import Button from "../../../components/button";
import minus from '../../../assets/minus.svg';
import plus from '../../../assets/plus.svg';

const { accessor, group } = createColumnHelper();

export const columns = [
  group({
    id: "@year",
    header: "Курс",
    columns: [
      accessor("year", {
        size: 100,
        header: "",
        cell: ({ row, subRows, getValue }) => (
          <div>
            {row.getCanExpand() ? (
              <Button
                {...{
                  onClick: row.getToggleExpandedHandler(),
                  style: { cursor: "pointer" },
                }}
              >
                {row.getIsExpanded() ? (
                  <img src={minus} alt="icon" style={{width: '1rem', height: '1rem'}}/>
                ) : (
                  <img src={plus} alt="icon" style={{width: '1rem', height: '1rem'}}/>
                )}
              </Button>
            ) : (
              null
            )}
            {getValue()}
          </div>
        ),
      }),
    ],
  }),
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
    id: "@id",
    header: "Id группы",
    columns: [
      accessor("id", {
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
