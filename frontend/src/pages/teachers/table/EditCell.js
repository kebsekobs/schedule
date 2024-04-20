import { Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { useDeleteTeacherMutation } from "../api/useDeleteTeacherMutation";
import { useState } from "react";
import EditModal from "../modals/EditModal.js";
import styles from "../../groups/table/table.module.css";

export function EditCell(props) {
  const deleteTeacherMutation = useDeleteTeacherMutation();
  const id = props.props.original.id;

  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const toggleEditModal = () => {
    setIsEditModalOpen(!isEditModalOpen);
  };
  function deleteTeacher() {
    if (window.confirm("Вы уверены, что хотите удалить преподавателя?"))
    deleteTeacherMutation.mutateAsync(id);
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
      <div className={styles["btns-wrapper"]} onClick={deleteTeacher}>
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
