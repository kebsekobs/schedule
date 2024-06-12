import React from "react";
import {
  flexRender,
  getCoreRowModel,
  useReactTable,
  getExpandedRowModel,
} from "@tanstack/react-table";
import { columns } from "./columns";
import styles from "../../shared/style/table.module.css";

const CoursesTable = ({ data, getRowCanExpand }) => {
  const table = useReactTable({
    columns,
    data,
    getRowCanExpand,
    getSubRows: (row) => row.subRows,
    getCoreRowModel: getCoreRowModel(),
    getExpandedRowModel: getExpandedRowModel(),
  });

  return (
    <table className={styles["table"]}>
      <thead>
        {table.getHeaderGroups().map((headerGroup) => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map((header) => (
              <th
                // style={{ width: "250px", border: "1px solid #a7a7a7" }} // тут нужен padding, но почему появляется второй пустой ряд заголовков??
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
          <>
            <tr className={styles["tr"]} key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <td className={styles["td"]} key={cell.id}>
                  <h3>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </h3>
                </td>
              ))}
            </tr>
          </>
        ))}
      </tbody>
    </table>
    // <table className={styles["table"]}>
    //   <thead>
    //     {table.getHeaderGroups().map((headerGroup) => (
    //       <tr className={styles["tr"]} key={headerGroup.id}>
    //         {headerGroup.headers.map((header) => (
    //           <th className={styles["th"]} key={header.id}>
    //             <h3>{header.column.columnDef.header}</h3>
    //             {/* <div
    //               className={`${styles["resizer"]} ${
    //                 header.column.getIsResizing() && styles["isResizing"]
    //               }`}
    //               onMouseDown={header.getResizeHandler}
    //               onTouchStart={header.getResizeHandler}
    //             ></div> */}
    //           </th>
    //         ))}
    //       </tr>
    //     ))}
    //   </thead>
    //   <tbody className={styles["table-body"]}>
    //     {table.getRowModel().rows.map((row) => (
    //       <tr key={row.id} className={styles["tr"]}>
    //         {row.getVisibleCells().map((cell) => (
    //           <td key={cell.id} className={styles["td"]}>
    //             <h3>
    //               {flexRender(cell.column.columnDef.cell, cell.getContext())}
    //             </h3>
    //           </td>
    //         ))}
    //       </tr>
    //     ))}
    //   </tbody>
    // </table>
  );
};

export default CoursesTable;
