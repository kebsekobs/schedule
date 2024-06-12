import { CellHelper } from "./CellHelpers";
import { createColumnHelper } from "@tanstack/react-table";
import { EditCell } from "./EditCell";
import Button from "../../../components/button";
import down from '../../../assets/down-arrow.svg';
import up from '../../../assets/up-arrow.svg';

const { accessor, group } = createColumnHelper();

export const columns = [
  group({
    id: "@year",
    header: "Курс",
    columns: [
      accessor("year", {
        size: 200,
        header: "",
        cell: ({ row, getValue }) => (
          <>
            {row.getCanExpand() ? (
              <span
                {...{
                  onClick: row.getToggleExpandedHandler(),
                  style: { cursor: "pointer" },
                }}
              >
                {row.getIsExpanded() ? (
                  <img src={up} alt="icon" style={{width: '2rem', height: '2rem', cursor: 'pointer', marginRight: '1rem'}}/>
                ) : (
                  <img src={down} alt="icon" style={{width: '2rem', height: '2rem', cursor: 'pointer', marginRight: '1rem'}}/>
                )}
              </span>
            ) : (
              null
            )}
            {getValue()}
          </>
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
        cell: ({ row }) => !row.getCanExpand() && <EditCell props={row} />,
      }),
    ],
  }),
];
