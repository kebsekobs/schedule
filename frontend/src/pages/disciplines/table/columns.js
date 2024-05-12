import {createColumnHelper} from '@tanstack/react-table';
import {CellHelper} from "./CellHelpers";
import {EditCell} from "./EditCell";

const {accessor, group} = createColumnHelper();

export const columns = [
    group({
        id: '@disciplinesId',
        header: 'Id',
        columns: [
            accessor('disciplinesId', {
                header: '',
                size: 400,
                cell: data =><div style={{ textAlign: 'center'}}>{data.getValue()}</div>,
            })
        ]
    }),
    group({
        id: '@name',
        header: 'Дисциплина',
        columns: [
            accessor('name', {
                header: '',
                size: 400,
                cell: data =><div style={{ textAlign: 'center'}}>{data.getValue()}</div>,
            })
        ]
    }),
    group({
        id: '@teachers',
        header: 'Преподаватель',
        columns: [
            accessor('teachers', {
                header: '',
                size: 400,
                cell: data =><div style={{ textAlign: 'center'}}>{data.getValue()}</div>,
            })
        ]
    }),
    group({
        id: '@hours',
        header: 'Часы/нед',
        columns: [
            accessor('hours', {
                header: '',
                size: 400,
                cell: data =><div style={{ textAlign: 'center'}}>{data.getValue()}</div>,
            })
        ]
    }),
    group({
        id: '@relatedGroupsId',
        header: 'Группы',
        columns: [
            accessor('relatedGroupsId', {
                header: '',
                size: 400,
                cell: data => CellHelper.renderRelatedGroups(data.row.original),
            })
        ]
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

