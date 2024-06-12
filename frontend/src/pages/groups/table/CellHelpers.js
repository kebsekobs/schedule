export const CellHelper = {
  renderGroupName: (row) => {
    const { groupId, magistracy } = row;

    if (isNaN(Math.floor(groupId / 100)))
      return;

    return `${Math.floor(groupId / 100)} Курс ${
      magistracy ? "Магистратуры" : "Бакалавриата"
    }`;
  },
};
