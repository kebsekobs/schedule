import React from "react";
import {
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { columns } from "./columns";
import styles from "../../groups/table/table.module.css";

const CoursesTable = ({ data }) => {
  const table = useReactTable({
    columns,
    data,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <table className={styles["table"]}>
      <thead>
        {table.getHeaderGroups().map((headerGroup) => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map((header) => (
              <th
                style={{ width: "250px", border: "2px solid #424242" }}
                key={header.id}
              >
                <h2>
                  {flexRender(
                    header.column.columnDef.header,
                    header.getContext()
                  )}
                </h2>
              </th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody className={styles["table-body"]}>
        {table.getRowModel().rows.map((row) => (
          <tr className={styles["tr"]} key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <td className={styles["td"]} key={cell.id}>
                <h3>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </h3>
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default CoursesTable;
