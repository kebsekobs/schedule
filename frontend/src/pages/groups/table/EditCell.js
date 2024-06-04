import { Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { useDeleteGroupMutation } from "../api/useDeleteGroupMutation";
import { useState } from "react";
import EditModal from "../modals/EditModal";
import styles from '../../shared/style/table.module.css';

export function EditCell(props) {
  const deleteGroupMutation = useDeleteGroupMutation();
  const id = props.props.original.id;

  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const toggleEditModal = () => {
    setIsEditModalOpen(!isEditModalOpen);
  };
  function deleteGroup() {
    if (window.confirm("Вы уверены, что хотите удалить группу?"))
      deleteGroupMutation.mutateAsync(id);
  }

  return (
    <div style={{padding: '5px'}}>
      <div
        className={styles["btns-wrapper"]}
        onClick={() => toggleEditModal(id)}
      >
        <p>Изменить</p>
        <Pencil1Icon />
      </div>
      <div className={styles["btns-wrapper"]} onClick={deleteGroup}>
        <p style={{ color: "var(--warning-color)" }}>Удалить</p>
        <TrashIcon />
      </div>
      {isEditModalOpen && (
        <EditModal
          toggleModal={toggleEditModal}
          isOpen={isEditModalOpen}
          id={id}
        />
      )}
    </div>
  );
}
