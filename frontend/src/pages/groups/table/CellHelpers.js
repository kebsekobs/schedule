export const CellHelper = {
    renderGroupName: (row) => {
        const {groupId, magistracy} = row;

        return `${Math.floor(groupId/100)} Курс ${magistracy ? 'Магистратуры' : 'Баклавариата'}`
    }
}